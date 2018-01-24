# helm-take-ownership
Transfer ownership of existing [kubernetes] release to [helm].

Let's say you have an existing [kubernetes] deployment that you want to migrate
to [helm]. [helm] currently doesn't have any functionality to migrate owneship
of those existing resources to itself. This plugin attempts to solve that
problem.

## How Does it Work?
[helm] v2 records information about a release in a configmap. When you try to
upgrade the release, [helm] uses this information to figure out what to do.
This plugin simply downloads the existing resource configurations, creates a
[helm] "Chart" with this information, and stores it in a configmap in the
format [helm] is expecting. You can then use [helm] to "upgrade" this release
using a Chart that you have built.

## How Does This Differ From Chartify?
Think of [helm-take-ownership] as a companion to [chartify]. [chartify] creates
a Chart from existing resources, but it doesn't migrate those resources to
[helm] control. [helm-take-ownership] can migrate those resources to [helm]'s
control. So, you migrate your resources with [helm-take-ownership], and then use
[chartify] to create a Chart from those resources. In the future, if you need
to make any changes to your Chart, you can use [helm] to upgrade the release
and everything should (hopefully) work beautifully, as if you used [helm] all
along.

## Installation
[helm-take-ownership] is a [helm] plugin. Installation is simple:

```bash
helm plugin install git@github.com:bmatcuk/helm-take-ownership.git
```

You'll need to have [go] and [glide] installed.

## Usage
[helm-take-ownership]'s operation is fairly simple. You tell it what resources
to include in the [helm] release using command line flags, plus any additional
[kubernetes] connection options (namespace, context, etc) and a name for the
release. For example:

```bash
helm own --deploy my-deployment --svc my-service -n stg my-release
```

The above example would instruct [helm-take-ownership] to associate a
deployment called `my-deployment` and a service called `my-service` in the
`stg` namespace to a [helm] release called `my-release`.

The naming of flags for [kubernetes] resources is the same as with `kubectl
get` (`--deployment`, `--service`, etc), including shorthands (`--deploy`,
`--svc`, etc). [kubernetes] connection options (such as namespace, context,
etc) are prepended with `k8s` (ex: `--k8s-namespace`, `--k8s-context`, etc).
If no [kubernetes] connection options are passed, it will use your default
configuration, the same way `kubectl` will.

Full usage is available from the command line:

```bash
helm own --help
```

## Some Notes
I built this tool for my own use. Currently, there are two things this plugin
will do that you might not want:

1. It will remove any `nodePort` values from services. In my setup, I'm relying
   on [kubernetes] giving me a random nodePort. When [helm-take-ownership]
   downloads the resources from [kubernetes], the random nodePort is in the
   data. If I then try to use [helm] to upgrade this release with a service
   definition that does not define a nodePort, [helm] removes the nodePort
   causing [kubernetes] to give me another port. This is undesirable.
2. It will remove any `selector` from a deployment spec. Similar to point #1:
   I'm relying on [kubernetes] giving me a default value here. When I do a
   [helm] upgrade, if I have the selector, [helm] removes it and [kubernetes]
   recreates it causing a new replicaset to be created. The existing replicaset
   is "lost" in the transition; orphaned and unmanaged.

I think, in the future, it might make sense to have some sort of "edit" option
that would allow a user to edit the Chart before it is installed to take care
of these very specific use cases.

## TODO
- [ ] command line switch to switch between [helm] configmaps and secrets
- [ ] command line switch to set the [helm] namespace (defaults to kube-system)
- [ ] allow a user to edit the Chart before installing it

[helm-take-ownership]: https://github.com/bmatcuk/helm-take-ownership
[kubernetes]: https://kubernetes.io/
[helm]: https://helm.sh/
[chartify]: https://github.com/appscode/chartify
[go]: https://golang.org/
[glide]: https://glide.sh/
