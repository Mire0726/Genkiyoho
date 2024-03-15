package weather

import (
	"fmt"
	"net/http"
	"time"

	"strings"
)


func CheckPollen(Pre string) bool {
	cityCode := GetCityCodeFromPrefecture(Pre)
	// 現在の日付をYYYYMMDD形式で取得
	currentDate := time.Now().Format("20060102")

	// APIのエンドポイントを構築
	apiURL := fmt.Sprintf("https://wxtech.weathernews.com/opendata/v1/pollen?citycode=%s&start=%s&end=%s", cityCode, currentDate, currentDate)

	fmt.Println("API Request URL:", apiURL)

	// HTTP GETリクエストを送信してみましょう
	response, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("Error making request:", err)
		return false
	}
	defer response.Body.Close()

	// レスポンスを処理します（ここではステータスコードのみ表示）
	fmt.Println("Response Status Code:", response.StatusCode)
	fmt.Println("Response :", response)
	return true
}

var PrefectureToCityCode = map[string]string{
    "Hokkaido": "01100", // 北海道札幌市
    "Aomori": "02201",
    "Iwate": "03201",
    "Miyagi": "04100",
    "Akita": "05201",
    "Yamagata": "06201",
    "Fukushima": "07201",
    "Ibaraki": "08201",
    "Tochigi": "09201",
    "Gunma": "10201",
    "Saitama": "11100",
    "Chiba": "12100",
    "Tokyo": "13100", // 東京都
    "Kanagawa": "14100",
    "Niigata": "15202",
    "Toyama": "16201",
    "Ishikawa": "17201",
    "Fukui": "18201",
    "Yamanashi": "19201",
    "Nagano": "20201",
    "Gifu": "21201",
    "Shizuoka": "22100",
    "Aichi": "23100",
    "Mie": "24201",
    "Shiga": "25201",
    "Kyoto": "26100",
    "Osaka": "27100", // 大阪府大阪市
    "Hyogo": "28100",
    "Nara": "29201",
    "Wakayama": "30201",
    "Tottori": "31201",
    "Shimane": "32201",
    "Okayama": "33201",
    "Hiroshima": "34201",
    "Yamaguchi": "35203",
    "Tokushima": "36201",
    "Kagawa": "37201",
    "Ehime": "38201",
    "Kochi": "39201",
    "Fukuoka": "40130", // 福岡県福岡市
    "Saga": "41201",
    "Nagasaki": "42201",
    "Kumamoto": "43201",
    "Oita": "44201",
    "Miyazaki": "45201",
    "Kagoshima": "46201",
    "Okinawa": "47201",
}
func GetCityCodeFromPrefecture(prefecture string) string {
	// 都道府県名を正規化（大文字に変換）
	prefecture = strings.ToUpper(prefecture)
	
	// マッピングから市町村コードを検索
	if code, exists := PrefectureToCityCode[prefecture]; exists {
		return code
	}
	
	return "" // 対応する市町村コードが見つからない場合
}