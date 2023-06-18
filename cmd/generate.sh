export PANTS_CONCURRENT=true

pants run cmd:bin -- refresh resources/plugins.toml resources/plugins.gen.toml
pants run  cmd:bin -- gen-readme resources/plugins.gen.toml resources/adhoc.toml resources/README.tmpl.md README.md
