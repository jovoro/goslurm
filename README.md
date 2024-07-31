# SLURM REST API SDK in Go
This repository contains the a set of type definitions (i.e. models) and endpoint calls implementing SLURM's REST API.

Please note the following discussion applies to SLURM version 22.05.9 whose documentation you can find [here][SLURM doc]. However,
it should be noted that what matters in this particular scenario is the version of the OpenAPI plugin this SDK targets which, for
us, is version 0.0.38. This plugin version is still shipped with SLURM 24.05 which as of the time of writing is the latest SLURM release.

In other words, this SDK should be usable 'as-is' with the latest SLURM version provided the `openapi/v0.0.38` plugin is configured
when starting `slurmrestd`.

Generating all these resources has been greatly simplified by the fact that SLURM implements the [OpenAPI v3.0.2 spec][OpenAPI v3.0.2 spec].
This means we can leverage the [OpenAPI Generator][OpenAPI Generator] tool to simply create the SDK as needed. We just need to provide the
OpenAPI spec file which we can get from a SLURM instance by simply querying the `/openapi` endpoint. This repository includes two OpenAPI
definitions:

1. **`api.yaml`**: This is the definition 'as-is' with a few minor tweaks needed to iron out some failing tests. These tweaks basically boil
                   down to making some types larger (i.e. `int -> int64`) so that some counters could later be unmarshalled.

1. **`api_ro.yaml`**: For this definition we have just removed all the `POST` and `DELETE` methods so that a 'read-only' client ws derived.
                      This SDK's main purpose was handling some of the complexity in terms of API interaction so that the SLURM [Telegraf][Telegraf]
                      plugin was smaller and easier to maintain. This is the definition this SDK's been derived from! The previous API definition
                      is simply provided as a convenience.

## Examples and usage
We provide a small example of how this client can be leveraged under the `cmd` directory. The idea is it should be you who imports this SDK in your
own Go project so that you can interact with a SLURM cluster. Importing the SDK can be accomplished with the regular `import` statement:

    import "github.com/pcolladosoto/go-slurm/v0.0.38"

## Authenticating with SLURM
The authentication of requests against the API leverages so called [JWT Tokens][JWT Tokens]. The basic idea is this token needs to be specified
when making the request through the `X-SLURM-USER-TOKEN` HTTP header together with the user the token belongs to through the `X-SLURM-USER-NAME`
HTTP header. However, the SDK exposes these in such a way that you needn't be concerned with meddling with the headers yourself. These authentication
parameters are passed through a [`context`][Go Context] as seen on the example under the `cmd` directory. Be sure to take a look at that.

Configuring the token-based authentication and generating the tokens themselves is covered on [SLURM's documentation][SLURM JWT Doc].

## Generating the SDK
As we said before generating the SDK is a matter of invoking OpenAPI generator. You'd need to install it first, which is very well explained on their
repository's [README.md][OpenAPI Generator]. Once that's done, you should be able to run the following from this directory (i.e. the one this `README.md`
is on):

    $ openapi-generator generate -g go -o oapigen -i api.yaml                          \
        --git-host github.com --git-user-id pcolladosoto --git-repo-id goslurm/oapigen \
        -c cfg.json

The `cfg.json` file contains the settings for the Go generator we invoke through the `-g go` flag. As always, please refer to the official documentation
on [general usage][OpenAPI Generator Usage], the [Go generator][Go Generator] and its [configuration][Go Generator Configuration].

Given we tend to forget these rather long commands we have condensed some of them in a `Makefile` you can invoke through `make(1)` if that's your jam.

## Further reading
It's worth taking a look at [SLURM's REST API documentation][SLURM REST Doc] as well as the [API reference][SLURM REST Reference] to have a clearer
overview of how everything works. Also, OpenAPI Generator defines a ton of markdown (i.e. `*.md`) files together with the SDK itself which you can
query at your leisure.

## Something missing?
If you find something is missing please do let us know or, even better, open a PR to add it! We're more than glad to accept contributions.

<!-- Get the links out of the text body to make things more readable! -->
[OpenAPI v3.0.2 spec]: https://spec.openapis.org/oas/v3.0.2
[SLURM doc]: https://slurm.schedmd.com/archive/slurm-22.05.9/documentation.html
[OpenAPI Generator]: https://github.com/OpenAPITools/openapi-generator
[Telegraf]: https://github.com/influxdata/telegraf
[OpenAPI Generator Usage]: https://openapi-generator.tech/docs/usage/
[Go Generator]: https://openapi-generator.tech/docs/generators/go/
[Go Generator Configuration]: https://openapi-generator.tech/docs/configuration/
[JWT Tokens]: https://en.wikipedia.org/wiki/JSON_Web_Token
[Go Context]: https://pkg.go.dev/context
[SLURM JWT Doc]: https://slurm.schedmd.com/archive/slurm-22.05.9/jwt.html
[SLURM Rest Doc]: https://slurm.schedmd.com/archive/slurm-22.05.9/rest.html
[SLURM Rest Reference]: https://slurm.schedmd.com/archive/slurm-22.05.9/rest_api.html
