// Since prism.js is written in JS, it doesn't have type definitions.
// We need to declare modules in a TypeScript d.ts file to avoid ts 7106 error: `Could not find a declaration file for module 'xxx/yyy'`.
//
declare module "prismjs/plugins/line-highlight/prism-line-highlight";
declare module "prismjs/plugins/line-numbers/prism-line-numbers";
declare module "prismjs/components/prism-protobuf";
