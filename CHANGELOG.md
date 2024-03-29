# Changelog

## v0.13 - 2022-10-20

## 💡 Enhancements 💡

- Support hostnames in Istio Virtual Service in `gateway-api`

## v0.12 - 2022-10-20

## 💡 Enhancements 💡

- Support hostnames in Istio Gateway in `gateway-api`

## v0.11 - 2022-10-19

## 💡 Enhancements 💡

- Update of istio mode in `gateway-api`

## v0.9 - 2022-09-29

## 💡 Enhancements 💡

- Add readyness probes for all frameworks
- Add appProtocol for services

## v0.8 - 2022-09-29

## 💡 Enhancements 💡

- Add support for Ingress in `gateway-api`
- Add support for restoring GatewayApi resources in  `gateway-api`
- Add `noop` mode in  `gateway-api`

## v0.7 - 2022-09-27

## 🛑 Breaking changes 🛑

- Envoy output doesn't create an Ingress resource by default 

## 🚀 New components 🚀

- Added experimental ArgoCD plugin

## 🧰 Bug fixes 🧰

- `ensure-sa` gets works again 

## 💡 Enhancements 💡

- Envoy output can enable/disable Ingress resource
- Add `generated` namespace to the resources

## v0.6 - 2022-09-24

## 🚀 New components 🚀

- First acceptable release. The following KRM functions are new:
  * create-atoms
  * ensure-sa
  * ensure-telemetry
  * gateway-api

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## Unreleased

## 🛑 Breaking changes 🛑

## 🚀 New components 🚀

## 🧰 Bug fixes 🧰

## 💡 Enhancements 💡
