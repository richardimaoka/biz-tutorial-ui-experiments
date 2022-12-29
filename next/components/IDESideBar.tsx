import { css } from "@emotion/react";

export const IDESideBar = (): JSX.Element => {
  const files = [
    { offset: 0, filename: "next" },
    { offset: 1, filename: "cache" },
    { offset: 1, filename: "server" },
    { offset: 1, filename: "static" },
    { offset: 1, filename: "build-manifest.json" },
    { offset: 1, filename: "package.json" },
    { offset: 1, filename: "react-loadable-manifest.json" },
    { offset: 1, filename: "trace" },
    { offset: 0, filename: "components" },
    { offset: 1, filename: "Terminal.tsx" },
    { offset: 0, filename: "node_module" },
    { offset: 0, filename: "pages" },
    { offset: 1, filename: "api" },
    { offset: 2, filename: "hello.ts" },
    { offset: 1, filename: "_app.tsx" },
    { offset: 1, filename: "_document.tsx" },
    { offset: 1, filename: "index.tsx" },
    { offset: 0, filename: "public" },
    { offset: 0, filename: "styles" },
    { offset: 1, filename: "global.css" },
    { offset: 1, filename: "Home.module.css" },
    { offset: 1, filename: ".babelrc" },
    { offset: 0, filename: ".eslintrc.json" },
    { offset: 0, filename: ".gitignore" },
    { offset: 0, filename: "next-env.d.ts" },
    { offset: 0, filename: "next.config.js" },
    { offset: 0, filename: "package-lock.json" },
    { offset: 0, filename: "package.json" },
    { offset: 0, filename: "README.md" },
    { offset: 0, filename: "tsconfig.json" },
  ];

  return (
    <div
      css={css`
        height: 340px;
        width: 160px;
        overflow: auto;
      `}
    >
      {files.map((elem, index) => (
        <div
          key={index}
          css={css`
            font-size: 12px;
            padding-left: ${8 * elem.offset}px;
          `}
        >
          {elem.filename}
        </div>
      ))}
    </div>
  );
};
