package handleData

import (
	"crawdata/database"
	"strings"
)

func defineTechnology(name string) (technology string) {
	if strings.Contains(name, "OLED") {
		technology = "OLED"
	} else {
		technology = "LED"
	}
	return
}
func defineResolution(name string) (resolution string) {
	if strings.Contains(name, "8K") {
		resolution = "8K"
	} else if strings.Contains(name, "4K") {
		resolution = "4K"
	} else {
		resolution = "Full HD"
	}
	return
}
func defineType(name string) (typee string) {
	if strings.Contains(name, "Google") {
		typee = "Google TV"
	} else if strings.Contains(name, "Android") {
		typee = "Android TV"
	} else {
		typee = "Smart TV"
	}
	return
}
func duplicateImg(input []string) (result []string) {
	result = append(result, input[0])
	result = append(result, input[0])
	result = append(result, input[0])
	return
}
func definePrices(prices []string, lengthSizeArr int) (result []string) {
	if len(prices) == 0 {
		for i := 0; i < lengthSizeArr; i++ {
			result = append(result, "50,000,000 VNĐ")
		}
	} else {
		result = prices
	}
	return
}
func HandleData(arr []database.Tivi) (result []database.Tivi) {
	linksdetail := []string{
		"https://www.sony.com.vn/electronics/tivi/z9j-series",
		"https://www.sony.com.vn/electronics/tivi/a90j-series",
		"https://www.sony.com.vn/electronics/tivi/a80j-series",
		"https://www.sony.com.vn/electronics/tivi/x95j-series",
		"https://www.sony.com.vn/electronics/tivi/x90j-series",
		"https://www.sony.com.vn/electronics/tivi/z8h-series",
		"https://www.sony.com.vn/electronics/tivi/a9g-series",
		"https://www.sony.com.vn/electronics/tivi/a9f-series",
		"https://www.sony.com.vn/electronics/tivi/z9f-series",
		"https://www.sony.com.vn/electronics/tivi/a9s-series",
		"https://www.sony.com.vn/electronics/tivi/a8h-series",
		"https://www.sony.com.vn/electronics/tivi/a8g-series",
		"https://www.sony.com.vn/electronics/tivi/a8f-series",
		"https://www.sony.com.vn/electronics/tivi/x95h-series",
		"https://www.sony.com.vn/electronics/tivi/x9500g-x9507g-series",
		"https://www.sony.com.vn/electronics/tivi/x90h-series",
		"https://www.sony.com.vn/electronics/tivi/x85j-series",
		"https://www.sony.com.vn/electronics/tivi/x8050h-series",
		"https://www.sony.com.vn/electronics/tivi/x85h-series",
		"https://www.sony.com.vn/electronics/tivi/x8500g-x8507g-x8577g-series",
		"https://www.sony.com.vn/electronics/tivi/x80j-series",
		"https://www.sony.com.vn/electronics/tivi/x80aj-series",
		"https://www.sony.com.vn/electronics/tivi/x80h-series",
		"https://www.sony.com.vn/electronics/tivi/x8000g-x8077g-series",
		"https://www.sony.com.vn/electronics/tivi/x75-series",
		"https://www.sony.com.vn/electronics/tivi/x74-x75a-series",
		"https://www.sony.com.vn/electronics/tivi/x7000g-x7007g-x7077g-series",
		"https://www.sony.com.vn/electronics/tivi/x9000f-series",
		"https://www.sony.com.vn/electronics/tivi/x8500f-series",
		"https://www.sony.com.vn/electronics/tivi/x8300f-series",
		"https://www.sony.com.vn/electronics/tivi/x75h-series",
		"https://www.sony.com.vn/electronics/tivi/x74h-series",
		"https://www.sony.com.vn/electronics/tivi/x7500f-series",
		"https://www.sony.com.vn/electronics/tivi/x7000f-series",
		"https://www.sony.com.vn/electronics/tivi/w800g-w802g-series",
		"https://www.sony.com.vn/electronics/tivi/w660g-series",
		"https://www.sony.com.vn/electronics/tivi/w610g-w617g-series",
		"https://www.sony.com.vn/electronics/tivi/w800f-series",
		"https://www.sony.com.vn/electronics/tivi/w660f-series",
		"https://www.sony.com.vn/electronics/tivi/w650d-series",
		"https://www.sony.com.vn/electronics/tivi/w600d-series",
	}
	imgs := [][]string{
		{

			"https://www.sony.com.vn/image/8ca05713e7748ff18c947e7a26dc6c9b?fmt=pjpeg&amp;wid=330&amp;hei=330&amp;bgcolor=F1F5F9&amp;bgc=F1F5F9",
		},
		{

			"https://www.sony.com.vn/image/ded064d5e0d54aa1db7666326421fec1?fmt=pjpeg&amp;wid=660&amp;hei=660&amp;bgcolor=F1F5F9&amp;bgc=F1F5F9",
		},
		{

			"https://www.sony.com.vn/image/928a63074f995fb9b8cf02dc991d8e15?fmt=pjpeg&amp;wid=660&amp;hei=660&amp;bgcolor=F1F5F9&amp;bgc=F1F5F9",
		},
		{

			"https://www.sony.com.vn/image/3044d380074338d5cebafbc767872c3a?fmt=pjpeg&amp;wid=660&amp;hei=660&amp;bgcolor=F1F5F9&amp;bgc=F1F5F9",
		},
		{

			"https://www.sony.com.vn/image/79f8e14a8cbf19eb2eae7065820c8855?fmt=pjpeg&amp;wid=660&amp;hei=660&amp;bgcolor=F1F5F9&amp;bgc=F1F5F9",
		},
		{

			"https://www.sony.com.vn/image/02a476ca1ffbb969f3de763539eb16d3?fmt=pjpeg&amp;wid=660&amp;hei=660&amp;bgcolor=F1F5F9&amp;bgc=F1F5F9",
		},
		{

			"https://www.sony.com.vn/image/7bedba2f19ee9130c0997a90cc2b976c?fmt=pjpeg&amp;wid=660&amp;hei=660&amp;bgcolor=F1F5F9&amp;bgc=F1F5F9",
		},
		{

			"https://www.sony.com.vn/image/d87e7af53de9581186e184bd55570697?fmt=pjpeg&amp;wid=660&amp;hei=660&amp;bgcolor=F1F5F9&amp;bgc=F1F5F9",
		},
		{

			"https://www.sony.com.vn/image/7318b1f4dc331b04ecfd0fd04a32b67c?fmt=pjpeg&amp;wid=660&amp;hei=660&amp;bgcolor=F1F5F9&amp;bgc=F1F5F9",
		},
		{

			"https://www.sony.com.vn/image/89859153b1890264f9a42c1b5aa4472a?fmt=pjpeg&amp;wid=660&amp;hei=660&amp;bgcolor=F1F5F9&amp;bgc=F1F5F9",
		},
		{

			"https://www.sony.com.vn/image/4443c62f0da678f768c11b44a3f045fc?fmt=pjpeg&amp;wid=660&amp;hei=660&amp;bgcolor=F1F5F9&amp;bgc=F1F5F9",
		},
		{

			"https://www.sony.com.vn/image/9b063b9c088d78ccbb4ce94ec4f82139?fmt=pjpeg&amp;wid=660&amp;hei=660&amp;bgcolor=F1F5F9&amp;bgc=F1F5F9",
		},
		{

			"https://www.sony.com.vn/image/bc620636d89ba27a1e8ec9fda59cd66a?fmt=pjpeg&amp;wid=660&amp;hei=660&amp;bgcolor=F1F5F9&amp;bgc=F1F5F9",
		},
		{

			"https://www.sony.com.vn/image/cc6df448e74241ae228061c6ed3bf182?fmt=pjpeg&amp;wid=660&amp;hei=660&amp;bgcolor=F1F5F9&amp;bgc=F1F5F9",
		},
		{

			"https://www.sony.com.vn/image/1b9c1142dd1d9bcf91ec4fd11af015da?fmt=pjpeg&amp;wid=660&amp;hei=660&amp;bgcolor=F1F5F9&amp;bgc=F1F5F9",
		},
		{

			"https://www.sony.com.vn/image/0d3b7fda823305ae524bac72fcb4b5cb?fmt=pjpeg&amp;wid=660&amp;hei=660&amp;bgcolor=F1F5F9&amp;bgc=F1F5F9",
		},
		{

			"https://www.sony.com.vn/image/ce286c6f75885a1fd8ba42eed8368596?fmt=pjpeg&amp;wid=660&amp;hei=660&amp;bgcolor=F1F5F9&amp;bgc=F1F5F9",
		},
		{

			"https://www.sony.com.vn/image/4dfd9ca48ba82b26b620b9049c099665?fmt=pjpeg&amp;wid=660&amp;hei=660&amp;bgcolor=F1F5F9&amp;bgc=F1F5F9",
		},
		{

			"https://www.sony.com.vn/image/0b8585e4d391720ae310dda542f3547d?fmt=pjpeg&amp;wid=660&amp;hei=660&amp;bgcolor=F1F5F9&amp;bgc=F1F5F9",
		},
		{

			"https://www.sony.com.vn/image/d79514a88ad8a1b185cbbe828384352c?fmt=pjpeg&amp;wid=660&amp;hei=660&amp;bgcolor=F1F5F9&amp;bgc=F1F5F9",
		},
		{

			"https://www.sony.com.vn/image/916e1ca5c3b64c3947acae853da7ad07?fmt=pjpeg&amp;wid=660&amp;hei=660&amp;bgcolor=F1F5F9&amp;bgc=F1F5F9",
		},
		{

			"https://www.sony.com.vn/image/d4f672c8c1b08401c3fb67cce747b7db?fmt=pjpeg&amp;wid=660&amp;hei=660&amp;bgcolor=F1F5F9&amp;bgc=F1F5F9",
		},
		{

			"https://www.sony.com.vn/image/4dfd9ca48ba82b26b620b9049c099665?fmt=pjpeg&amp;wid=660&amp;hei=660&amp;bgcolor=F1F5F9&amp;bgc=F1F5F9",
		},
		{

			"https://www.sony.com.vn/image/7860395db486205652e719f797803ec0?fmt=pjpeg&amp;wid=660&amp;hei=660&amp;bgcolor=F1F5F9&amp;bgc=F1F5F9",
		},
		{

			"https://www.sony.com.vn/image/7a5f71c31eace4488a85d3156b897d87?fmt=pjpeg&amp;wid=660&amp;hei=660&amp;bgcolor=F1F5F9&amp;bgc=F1F5F9",
		},
		{

			"https://www.sony.com.vn/image/d87192de6310f706f034bf33847fd8ab?fmt=pjpeg&amp;wid=660&amp;hei=660&amp;bgcolor=F1F5F9&amp;bgc=F1F5F9",
		},
		{

			"https://www.sony.com.vn/image/3c64c516ffdbd96ed39b83bd663596cd?fmt=pjpeg&amp;wid=660&amp;hei=660&amp;bgcolor=F1F5F9&amp;bgc=F1F5F9",
		},
		{

			"https://www.sony.com.vn/image/969b68cc1af4ace6a9a7530233bda21e?fmt=pjpeg&amp;wid=660&amp;hei=660&amp;bgcolor=F1F5F9&amp;bgc=F1F5F9",
		},
		{

			"https://www.sony.com.vn/image/d974fdb6e6886298ee5fc07584c93a47?fmt=pjpeg&amp;wid=660&amp;hei=660&amp;bgcolor=F1F5F9&amp;bgc=F1F5F9",
		},
		{

			"https://www.sony.com.vn/image/13694b0ddd3881e81ee7ca9611d90667?fmt=pjpeg&amp;wid=660&amp;hei=660&amp;bgcolor=F1F5F9&amp;bgc=F1F5F9",
		},
		{

			"https://www.sony.com.vn/image/01b495653bff0a0476302affff4ef731?fmt=pjpeg&amp;wid=660&amp;hei=660&amp;bgcolor=F1F5F9&amp;bgc=F1F5F9",
		},
		{

			"https://www.sony.com.vn/image/b0fdbd2af51a5b99deb7ff2a82dc36de?fmt=pjpeg&amp;wid=660&amp;hei=660&amp;bgcolor=F1F5F9&amp;bgc=F1F5F9",
		},
		{

			"https://www.sony.com.vn/image/d553a22107969f70107c7c417f979a7d?fmt=pjpeg&amp;wid=660&amp;hei=660&amp;bgcolor=F1F5F9&amp;bgc=F1F5F9",
		},
		{

			"https://www.sony.com.vn/image/88cea9de40220b1104e45cce4834ea04?fmt=pjpeg&amp;wid=660&amp;hei=660&amp;bgcolor=F1F5F9&amp;bgc=F1F5F9",
		},
		{

			"https://www.sony.com.vn/image/7ac1d36b8cf1686e98ef46d06e6c557e?fmt=pjpeg&amp;wid=660&amp;hei=660&amp;bgcolor=F1F5F9&amp;bgc=F1F5F9",
		},
		{

			"https://www.sony.com.vn/image/aa7c42f6c52b7b76d9f9fd5698a84fa4?fmt=pjpeg&amp;wid=330&amp;hei=330&amp;bgcolor=F1F5F9&amp;bgc=F1F5F9",
		},
		{

			"https://www.sony.com.vn/image/13463223bd5ed0d2328185e89b027d87?fmt=pjpeg&amp;wid=330&amp;hei=330&amp;bgcolor=F1F5F9&amp;bgc=F1F5F9",
		},
		{

			"https://www.sony.com.vn/image/e015d1beeeb5a092213452cee9bb7268?fmt=pjpeg&amp;wid=330&amp;hei=330&amp;bgcolor=F1F5F9&amp;bgc=F1F5F9",
		},
		{

			"https://www.sony.com.vn/image/23b090be2b2ac87c3b495afc2f531928?fmt=pjpeg&amp;wid=660&amp;hei=660&amp;bgcolor=F1F5F9&amp;bgc=F1F5F9",
		},
		{

			"https://www.sony.com.vn/image/7330e6d37c178d8930475bcc68628ca3?fmt=pjpeg&amp;wid=660&amp;hei=660&amp;bgcolor=F1F5F9&amp;bgc=F1F5F9",
		},
		{

			"https://www.sony.com.vn/image/7330e6d37c178d8930475bcc68628ca3?fmt=pjpeg&amp;wid=660&amp;hei=660&amp;bgcolor=F1F5F9&amp;bgc=F1F5F9",
		},
		{

			"https://www.sony.com.vn/image/3c97d9b99fa516fa7a415c09888cbbe8?fmt=pjpeg&amp;wid=660&amp;hei=660&amp;bgcolor=F1F5F9&amp;bgc=F1F5F9",
		},
	}
	for i, value := range arr {
		if i == 0 {
			continue
		}
		result = append(result, database.Tivi{
			ID:          int64(i),
			Name:        value.Name,
			Technology:  defineTechnology(value.Name),
			Resolution:  defineResolution(value.Name),
			Type:        defineType(value.Name),
			Imgs:        duplicateImg(imgs[i]),
			Description: value.Description,
			Sizes:       value.Sizes,
			Prices:      definePrices(value.Prices, len(value.Sizes)),
			LinkDetail:  linksdetail[i-1],
		})
	}
	return
}

// {41 W60D | LED | HD Ready/Full HD | Smart TV    []
// [X-Reality PRO cho chất lượng HD tốt hơn
// Truy cập YouTube và nhiều dịch vụ khác với Wi-Fi tích hợp
// Thiết kế thanh mảnh, phù hợp với phòng khách
// ]
// [32” (80 cm)]
// [6,590,000 VNĐ]}

// type Tivi struct {
// 	ID          int64    `json:"id"`
// 	Name        string   `json:"name"`
// 	Technology 		string   `json:"technology"`
// 	Resolution 		string   `json:"recursion"`
// 	Type 		string   `json:"type"`
// 	Imgs         []string   `json:"imgs"`
// 	Description []string `json:"description"`
// 	Sizes       []string `json:"sizes"`
// 	Prices      []string `json:"prices"`
// }
