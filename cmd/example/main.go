package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	slurmClient "github.com/jovoro/goslurm/v0039"
)

var (
	slurmUser  string
	slurmToken string
	host       string
	port       int
	scheme     string
)

func init() {
	flag.StringVar(&slurmUser, "slurmUser", "foo", "SLURM username for authentication")
	flag.StringVar(&slurmToken, "slurmToken", "topSecret", "SLURM JWT token for authentication")
	flag.StringVar(&host, "host", "localhost", "Host where the slurmrestd daemon is running")
	flag.IntVar(&port, "port", 6820, "Port where the slurmrestd daemon is listening")
	flag.StringVar(&scheme, "scheme", "http", "Scheme to use, one of http or https")
}

func main() {
	// Parse the provided flags
	flag.Parse()

	// If the authentication parameters remain untouched simply tell the user...
	if slurmUser == "foo" || slurmToken == "topSecret" {
		fmt.Printf("it looks like you haven't configured the authentication parameters... Maybe run %s --help?\n", os.Args[0])
	}

	// If the scheme is neither HTTP nor HTTPS fail
	scheme = strings.ToLower(scheme)
	if scheme != "http" && scheme != "https" {
		fmt.Printf("the provided scheme %s is neither http not https...\n", scheme)
		os.Exit(-1)
	}

	// Time to create the client and provide the necessary configuration options
	configuration := slurmClient.NewConfiguration()
	configuration.Host = fmt.Sprintf("%s:%d", host, port)
	configuration.Scheme = scheme

	// Let's create the context in which we embed the authentication parameters. This auth
	// context will later be passed to the different requests. Be sure to read up more on
	// the context package (which is part of the standard library) to get a feel for what
	// it achieves. The blog post at https://go.dev/blog/context is also a great read!
	auth := context.WithValue(
		context.Background(),
		slurmClient.ContextAPIKeys,
		map[string]slurmClient.APIKey{
			"user":  {Key: slurmUser},
			"token": {Key: slurmToken},
		},
	)

	// Time to create the client! We can now begin making requests
	apiClient := slurmClient.NewAPIClient(configuration)

	fmt.Printf("time to make a ping request!\n")

	// Let's call the ping API endpoint. You can find information on it on the following link
	// https://slurm.schedmd.com/archive/slurm-22.05.9/rest_api.html#slurmV0038Ping
	resp, httpRes, err := apiClient.SlurmAPI.SlurmV0038Ping(auth).Execute()
	if err != nil {
		fmt.Printf("error making the ping request: %v\n", err)

		// As the request returned an error we'll see the response is not populated...
		if resp == nil {
			fmt.Printf("the response is nil... Let's try to decode it ourselves!\n")
		}

		// Let's read out the entire body and parse it like we would any other JSON object.
		body, err := io.ReadAll(httpRes.Body)
		if err != nil {
			fmt.Printf("error getting the raw HTTP body: %v\n", err)
			os.Exit(-1)
		}
		respErrors := slurmClient.V0038Errors{}
		if err := json.Unmarshal(body, &respErrors); err != nil {
			fmt.Printf("error unmarshalling the errors: %v\n", err)
			os.Exit(-1)
		}

		// Now that we've got the parsed errors let's print them out!
		for i, respError := range respErrors.Errors {
			fmt.Printf("\t.errors.%d.error: %s\n", i, *respError.Error)
			fmt.Printf("\t.errors.%d.error_number: %d\n", i, *respError.ErrorNumber)
			fmt.Printf("\t.errors.%d.description: %s\n", i, *respError.Description)
			fmt.Printf("\t.errors.%d.source: %s\n", i, *respError.Source)
		}
		os.Exit(-1)
	}

	// We can access the raw HTTP response of type *http.Response. We already saw it
	// when decoding the error above though!
	fmt.Printf("\nthe response's status code is: %s\n", httpRes.Status)

	// And we can also access the parsed body of the response through resp. Given most of the
	// members of the reply are pointers we can use some helper methods to check whether the
	// data pointed to is actually populated or if on the contrary it's simply the nil value
	// for the underlying type. That way we can 'unroll' the data as much as we need!
	meta, ok := resp.GetMetaOk()
	if !ok {
		fmt.Printf("the meta field is not populated...\n")
		os.Exit(-1)
	}
	fmt.Printf("\nprinting out the meta field...\n")
	printMeta(meta)

	// Let's do the same thing for the pings field
	pings, ok := resp.GetPingsOk()
	if !ok {
		fmt.Printf("the pings field is not populated...\n")
		os.Exit(-1)
	}
	fmt.Printf("\nprinting out the pings field...\n")
	printPings(pings)

	fmt.Printf("\nthat's all folks!\n")
	os.Exit(0)
}

func printMeta(meta *slurmClient.V0038Meta) {
	slurmInfo, ok := meta.GetSlurmOk()
	if !ok {
		fmt.Printf("\tthe .meta.slurm field is not populated...\n")
		return
	}

	// If the fields themselves are empty are empty we'll just show these empty strings not to make the function too long...
	fmt.Printf("\t.meta.slurm.release: %s\n", *slurmInfo.Release)

	slurmVersion, ok := slurmInfo.GetVersionOk()
	if !ok {
		fmt.Printf("\tthe .meta.slurm.version field is not populated...\n")
		return
	}
	fmt.Printf("\t.meta.slurm.version.major: %d\n", *slurmVersion.Major)
	fmt.Printf("\t.meta.slurm.version.minor: %d\n", *slurmVersion.Minor)
	fmt.Printf("\t.meta.slurm.version.micro: %d\n", *slurmVersion.Micro)

	pluginInfo, ok := meta.GetPluginOk()
	if !ok {
		fmt.Printf("\tthe .meta.plugin field is not populated...\n")
		return
	}

	fmt.Printf("\t.meta.plugin.name: %s\n", *pluginInfo.Name)
	fmt.Printf("\t.meta.plugin.type: %s\n", *pluginInfo.Type)
}

func printPings(pings []slurmClient.V0038Ping) {
	// Given we're accessing the fields directly we're okay with printing out the
	// nil fields for both strings and integers...
	for i, ping := range pings {
		fmt.Printf("\t.pings.%d.mode: %s\n", i, *ping.Mode)
		fmt.Printf("\t.pings.%d.hostname: %s\n", i, *ping.Hostname)
		fmt.Printf("\t.pings.%d.ping: %s\n", i, *ping.Ping)
		fmt.Printf("\t.pings.%d.status: %d\n", i, *ping.Status)
	}
}
