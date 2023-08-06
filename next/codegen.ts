import type { CodegenConfig } from "@graphql-codegen/cli";

const config: CodegenConfig = {
  overwrite: true,
  schema: "../gqlgen/graph/schema.gql",
  documents: ["app/**/*.tsx"],
  generates: {
    "libs/gql/": {
      preset: "client",
      plugins: [],
      config: {
        // avoidOptionals: true,
        nonOptionalTypename: true, //__typename should always exist
        enumsAsTypes: true,
        dedupeFragments: true, //GraphQL】 There can be only one fragment named XxxFragment error - https://zenn.dev/bicstone/articles/graphql-dedupe-fragments, https://github.com/dotansimha/graphql-code-generator/issues/8670
      },
    },
  },
  watch: ["app/**/*.tsx", "../gqlgen/graphq/schema.gql"],
  hooks: { afterOneFileWrite: ["npx prettier --write"] },
};

export default config;
