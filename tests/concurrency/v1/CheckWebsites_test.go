package v1

type WebSiteChecker func(string) bool

func CheckWebsites(wc WebSiteChecker, urls []string) map[string]bool {
	result := make(map[string]bool)

	for _, url := range urls {
		go func() { // bu bir anonymus fonksiyon. Sonundaki () tanimlandiklari anda calisabilmesini sagliyor
			result[url] = wc(url)
		}()
	}
	return result
}
