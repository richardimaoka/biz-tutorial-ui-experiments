// Since prism.js is written in JS, it doesn't have type definitions.
// We need to declare modules in a TypeScript d.ts file to avoid ts 7106 error: `Could not find a declaration file for module 'xxx/yyy'`.
//
// to fix ts 7106 error in `import "prismjs/components/prism-protobuf";`
declare module "prismjs/components/prism-protobuf";
