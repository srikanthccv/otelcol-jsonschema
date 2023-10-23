
_There is a [proposal](https://docs.google.com/document/d/15SMsJAdE1BHAmwf8cLXBvsHDrezXOHniaRZafjO-5e4/edit) and work-in-progress PR for generating the component config from the YAML or JSON Schema [here](https://github.com/open-telemetry/opentelemetry-collector-contrib/pull/27003). I am not working on perfecting this, but it was a fun and good learning experience_


# OpenTelemetry Collector JSON Schema

This repository contains the [JSON Schema](https://json-schema.org/) for the OTEL Collector YAML configuration file.

## Why

- See the status, supported distributions and overview documentation of the component without leaving your editor.
- Get auto-completion and validation of the configuration file.
- And all other benefits of using a JSON Schema.

## Usage

The schema is available at https://raw.githubusercontent.com/srikanthccv/otelcol-jsonschema/main/schema.json. Many popular editors and IDEs support YAML validation against a JSON Schema. For example, in VS Code, you can add the following to your `settings.json`:

```json
"yaml.schemas": {
    "https://raw.githubusercontent.com/srikanthccv/otelcol-jsonschema/main/schema.json": "otel-collector-config.yaml"
}
```

See https://github.com/redhat-developer/yaml-language-server#clients for other choices.

## TODO

- [ ] Autogenerate the /components/...
- [ ] null or reference for the "type" field using "oneOf"
- [ ] Support various distributions of the collector
- [ ] Improve the validation of "service"
- [ ] Workaround for tricky components [filelog, hostmetrics, etc...]

Any contributions are welcome. Please open an issue or a pull request.

## Credits

To authors of https://github.com/invopop/jsonschema, which is modified to make it work with the OTEL Collector Config types.