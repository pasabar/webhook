package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/aiteung/atapi"
	"github.com/aiteung/atmessage"
	"github.com/aiteung/module/model"
	"github.com/pasabar/webhook"
	"github.com/pasabar/webhook/functions"
	"github.com/whatsauth/wa"
)

func PostBalasan(w http.ResponseWriter, r *http.Request) {
	var msg model.IteungMessage
	var resp atmessage.Response
	json.NewDecoder(r.Body).Decode(&msg)
	link := "https://medium.com/@ryaasishlah123/langkah-langkah-implementasi-whatsauth-free-integrasi-2fa-otp-dan-notifikasi-melalui-whatsapp-f462fcb7ea25"
	if r.Header.Get("Secret") == os.Getenv("SECRET") {
		if msg.Message == "loc" || msg.Message == "Loc" || msg.Message == "lokasi" || msg.LiveLoc {
			location, err := functions.ReverseGeocode(msg.Latitude, msg.Longitude)
			if err != nil {
				// Handle the error (e.g., log it) and set a default location name
				location = "Unknown Location"
			}

			reply := fmt.Sprintf("Sampurasun kamu pasti lagi di %s "+
				"\n Koordinatenya : %s - %s"+
				"\n Cara Penggunaan WhatsAuth Ada di link dibawah ini"+
				"yaa kak %s\n", location,
				strconv.Itoa(int(msg.Longitude)), strconv.Itoa(int(msg.Latitude)), link)
			dt := &wa.TextMessage{
				To:       msg.Phone_number,
				IsGroup:  false,
				Messages: reply,
			}
			resp, _ = atapi.PostStructWithToken[atmessage.Response]("Token", os.Getenv("TOKEN"), dt, "https://api.wa.my.id/api/send/message/text")

		} else if msg.Message == "gandeng" || msg.Message == "Anjing" || msg.Message == "goblok" {
			dt := &wa.TextMessage{
				To:       msg.Phone_number,
				IsGroup:  false,
				Messages: fmt.Sprintf("Kalemken atuh %s jangan toksis kitu, bisi kakenca", msg.Alias_name),
			}
			resp, _ = atapi.PostStructWithToken[atmessage.Response]("Token", os.Getenv("TOKEN"), dt, "https://api.wa.my.id/api/send/message/text")

		} else if msg.Message == "cantik" || msg.Message == "ganteng" || msg.Message == "cakep" {
			dt := &wa.TextMessage{
				To:       msg.Phone_number,
				IsGroup:  false,
				Messages: fmt.Sprintf("Hatur nuhun %s kamu jugaa segalanya banget", msg.Alias_name),
			}
			resp, _ = atapi.PostStructWithToken[atmessage.Response]("Token", os.Getenv("TOKEN"), dt, "https://api.wa.my.id/api/send/message/text")

		} else if strings.Contains(msg.Message, "login") {
			//login username test password testcihuy
			messages := strings.Split(msg.Message, " ")
			email := messages[2]
			password := messages[len(messages)-1]
			dt := &webhook.Logindata{
				Email:    email,
				Password: password,
			}
			res, _ := atapi.PostStructWithToken[atmessage.Response]("Token", os.Getenv("TOKEN"), dt, "https://asia-southeast2-pasabar.cloudfunctions.net/Admin-Login")
			dat := &wa.TextMessage{
				To:       msg.Phone_number,
				IsGroup:  false,
				Messages: res.Response,
			}
			resp, _ = atapi.PostStructWithToken[atmessage.Response]("Token", os.Getenv("TOKEN"), dat, "https://api.wa.my.id/api/send/message/text")
		} else {
			randm := []string{
				"Sampurasun brader " + msg.Alias_name + "\n ryaas & fahad lagi gaadaa \n aku pasabarbot salam kenall yaaaa \n Cara penggunaan WhatsAuth ada di link berikut ini ya kak...\n" + link,
				"OI jangan SPAAM berisik tau pasabar lagi tidur",
				"Kamu ganteng tau, tapi masi gantengan iqbal",
				"Ihhh kamu cantik banget, tapi masih cantikan mamahnya aku",
				"bro, Siapa nama bapak kamu?",
				"Disini gaboleh marah marah, pasti abis losestreak ya?",
				"Siapa disini yang nama bapaknya MUSALODONS?",
				"ALAGADAY TU THESKAY TU THEMOON",
				"Hereuy hehehe",
			}
			dt := &wa.TextMessage{
				To:       msg.Phone_number,
				IsGroup:  false,
				Messages: functions.GetRandomString(randm),
			}
			resp, _ = atapi.PostStructWithToken[atmessage.Response]("Token", os.Getenv("TOKEN"), dt, "https://api.wa.my.id/api/send/message/text")
		}
	} else {
		resp.Response = "Secret Salah"
	}
	fmt.Fprintf(w, resp.Response)
}
