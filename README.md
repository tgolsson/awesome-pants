# Awesome Pants

This repository is a collection of great Pants tools, guides, and plugins.

It is driven by code generation; in order to be presented both as a website and as a README.

Each list is contained in a separate file in `/resources/`: `plugins` for plugins, and so on. Just follow the existing pattern and then run `pants generate` to regenerate all other lists! The website itself is rebuilt by CI and published, so you just need to ensure the README is regenerated.

## Plugins

- [pants-backend-bitwarden](https://pypi.org/project/pants-backend-bitwarden/) - Version: 0.2.0  
  A Bitwarden plugin for the Pants build system

- [pants-backend-k8s](https://pypi.org/project/pants-backend-k8s/) - Version: 0.2.0  
  A Kubernetes plugin for the Pants build system

- [pants-backend-kustomize](https://pypi.org/project/pants-backend-kustomize/) - Version: 0.2.0  
  A Kustomize plugin for the Pants build system

- [pants-backend-mdbook](https://pypi.org/project/pants-backend-mdbook/) - Version: 0.2.0  
  A  MdBook documentation builder plugin for the Pants buildsystem.

- [pants-backend-oci](https://pypi.org/project/pants-backend-oci/) - Version: 0.4.0  
  An OCI plugin for the Pants build system

- [pants-backend-secrets](https://pypi.org/project/pants-backend-secrets/) - Version: 0.2.0  
  A secrets plumbing plugin for the Pants build system


## Recipes

- [SvelteKit building](https://gist.github.com/sureshjoshi/98fb09f2a340f7c1dad270c4887865a0#file-build-pants-sveltekit) by [Suresh Joshi](https://github.com/sureshjoshi)  
  Driving Sveltekit builds with pNPM and node, and generating an Azure archive.

- [The Pants adhoc examples repository](https://github.com/pantsbuild/example-adhoc) by [The Pantsbuild team](https://github.com/pantsbuild)  
  A collection of recipes maintained by the Pantsbuild organization.



