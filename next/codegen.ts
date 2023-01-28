import type { CodegenConfig } from "@graphql-codegen/cli";

const config: CodegenConfig = {
  overwrite: true,
  schema: "../gqlgen/graph/schema.gql",
  documents: ["pages/**/*.tsx", "components/**/*.tsx"],
  generates: {
    "libs/gql/": {
      preset: "client",
      plugins: [],
      config: {
        // avoidOptionals: true,
        nonOptionalTypename: true, //__typename should always exist
        enumsAsTypes: true,
      },
    },
  },
  watch: [
    "pages/**/*.tsx",
    "components/**/*.tsx",
    "../gqlgen/graphq/schema.gql",
  ],
  hooks: { afterOneFileWrite: ["npx prettier --write"] },
};

export default config;
