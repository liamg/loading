package bar

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
)

func Download(uri string, w io.Writer) error {

	parsed, err := url.Parse(uri)
	if err != nil {
		return err
	}

	// create the bar
	loadingBar := New(
		OptionWithStatsFuncs(
			StatsPercentComplete,
			StatsBytesComplete,
			StatsBytesRate,
			StatsTimeRemaining,
		),
		OptionWithLabel(fmt.Sprintf("Downloading %s...", path.Base(parsed.Path))),
	)

	// download a file
	resp, err := http.DefaultClient.Get(uri)
	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() }()

	// tell the loading bar how big the total download is
	loadingBar.SetTotalInt64(resp.ContentLength)

	// download the data to the temporary file, and update the loading bar along the way
	if _, err := io.Copy(io.MultiWriter(w, loadingBar), resp.Body); err != nil {
		loadingBar.Finish()
		return err
	}

	return nil
}
