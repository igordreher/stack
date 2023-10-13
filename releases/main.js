import fs from 'fs'
import YAML from 'yaml'

// Get Version from ./version/v1.1.0.yml
const VERSION = process.env.VERSION || "v1.1.0"
const dirVersion = `./versions/${VERSION}`
const versionFile = fs.readFileSync(`${dirVersion}/main.yaml`, 'utf8')
const CONTENT = YAML.parse(versionFile)

// Generate openapi-merge.json
const openapiFile = fs.readFileSync(`./templates/openapi-merge.json`, 'utf8')
const openapiConfig = openapiFile
    .replace('AUTH_VERSION', CONTENT.components.auth)
    .replace('LEDGER_VERSION', CONTENT.components.ledger)
    .replace('PAYMENTS_VERSION', CONTENT.components.payments)
    .replace('SEARCH_VERSION', CONTENT.components.search)
    .replace('ORCHESTRATION_VERSION', CONTENT.components.orchestration)
    .replace('WALLETS_VERSION', CONTENT.components.wallets)
    .replace('WEBHOOKS_VERSION', CONTENT.components.webhooks)
    .replace('STARGATE_VERSION', CONTENT.components.stargate)
    .replace('GATEWAY_VERSION', CONTENT.components.gateway)
    .replace('OUTPUT_FILE', `./../${dirVersion}/openapi.json`)

// Write file in ./build/openapi-merge.json
fs.writeFileSync(`./build/openapi-merge.json`, openapiConfig)