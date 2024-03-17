package weather

import (
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"strings"
)


func CheckPollen(Pre string) int{
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
		return 0
	}

	defer response.Body.Close()

    r := csv.NewReader(response.Body)
    // 最大花粉飛散量を記録する変数
	maxPollen := 0

	// CSVデータを1行ずつ読み込む
	for {
		record, err := r.Read()
		if err == io.EOF {
			break // ファイルの終わりに達した
		}
		if err != nil {
			fmt.Println("Error reading CSV:", err)
			return 0
		}

		// 花粉飛散量を取得して整数に変換
		pollen, err := strconv.Atoi(record[len(record)-1])
		if err != nil {
			fmt.Println("Error converting pollen count:", err)
			continue
		}

		// 最大値を更新
		if pollen > maxPollen {
			maxPollen = pollen
		}
	}

	fmt.Println("Maximum pollen count for the day:", maxPollen)
    return maxPollen
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