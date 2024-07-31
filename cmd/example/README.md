# An example on how to use the client
This directory contains a file showcasing how one might go about using the client implemented in the repository.

The example makes a [`/ping`](https://slurm.schedmd.com/archive/slurm-22.05.9/rest_api.html#slurmV0038Ping) request
to a SLURM daemon to then print out all the received information. In case of an error, all the available data will
be printed as well.

The code itself is heavily commented and even some links are provided for further reading. We encourage you to read it.

## Compiling the example
Compilation is just a matter of running:

    $ go build

This has to be run of course from this directory. Now, isn't Go wonderfully simple?

## Running the example
The example implements a rudimentary CLI based on the standard `flag` package. Be sure to run `./example --help`
to get a list of available flags. Every bit of information needed for running the program is to be provided
through these flags. This includes the SLURM user name and JWT token as well as the host where the `slurmrestd`
daemon is running. The port `slurmrestd` is bound to and the URL scheme (i.e. `http` or `https`) can also be
chosen. The default scheme is `http`.

For the sake of comfort, we usually define our SLURM credentials through variables to make invoking programs
a bit easier. We populate them on a regular shell with:

    $ SLURM_USER=foo

and:

    $ SLURM_JWT=topSecret

so that we can finally run:

    $  ./example --host yourMachine --slurmUser "$SLURM_USER" --slurmToken "$SLURM_JWT"

You can also, of course, pass them explicitly: up to you!

Anyway, when you run the program you'll see some information printed on screen. It's actually kind of lame, but
the important thing is it showcases how everything works.
