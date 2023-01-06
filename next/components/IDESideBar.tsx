import { css } from "@emotion/react";
import { FileNameComponent } from "./IDE/filetree/FileNameComponent";

interface IDESideBarProps {
  height: number;
}

const files = [
  { offset: 0, __typename: "directory", filename: "next" },
  { offset: 1, __typename: "directory", filename: "cache" },
  { offset: 1, __typename: "directory", filename: "server" },
  { offset: 1, __typename: "directory", filename: "static" },
  { offset: 1, __typename: "file", filename: "build-manifest.json" },
  { offset: 1, __typename: "file", filename: "package.json" },
  {
    offset: 1,
    __typename: "file",
    filename: "react-loadable-manifest.json",
  },
  { offset: 1, __typename: "directory", filename: "trace" },
  { offset: 0, __typename: "directory", filename: "components" },
  { offset: 1, __typename: "file", filename: "Terminal.tsx" },
  { offset: 0, __typename: "directory", filename: "node_module" },
  { offset: 0, __typename: "directory", filename: "pages" },
  { offset: 1, __typename: "directory", filename: "api" },
  { offset: 2, __typename: "file", filename: "hello.ts" },
  { offset: 1, __typename: "file", filename: "_app.tsx" },
  { offset: 1, __typename: "file", filename: "_document.tsx" },
  { offset: 1, __typename: "file", filename: "index.tsx" },
  { offset: 0, __typename: "directory", filename: "public" },
  { offset: 0, __typename: "directory", filename: "styles" },
  { offset: 1, __typename: "file", filename: "global.css" },
  { offset: 1, __typename: "file", filename: "Home.module.css" },
  { offset: 1, __typename: "directory", filename: ".babelrc" },
  { offset: 0, __typename: "file", filename: ".eslintrc.json" },
  { offset: 0, __typename: "directory", filename: ".gitignore" },
  { offset: 0, __typename: "file", filename: "next-env.d.ts" },
  { offset: 0, __typename: "file", filename: "next.config.js" },
  { offset: 0, __typename: "file", filename: "package-lock.json" },
  { offset: 0, __typename: "file", filename: "package.json" },
  { offset: 0, __typename: "file", filename: "README.md" },
  { offset: 0, __typename: "file", filename: "tsconfig.json" },
];

export const IDESideBar = ({ height }: IDESideBarProps): JSX.Element => {
  return (
    <div
      css={css`
        height: ${height}px;
        max-width: 160px;
        overflow: auto;
        ::-webkit-scrollbar {
          width: 5px;
          height: 5px;
          background-color: #252526; /* or add it to the track */
        }
        ::-webkit-scrollbar-thumb {
          background: #a0a0a0;
          border-radius: 5px;
        }
      `}
    >
      <div
        css={css`
          width: fit-content;
        `}
      >
        {files.map((elem) => (
          <FileNameComponent
            key={elem.filename}
            type={elem.__typename}
            offset={elem.offset}
            name={elem.filename}
          />
        ))}
      </div>
    </div>
  );
};
