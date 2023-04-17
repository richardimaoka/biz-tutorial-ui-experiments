import { css } from "@emotion/react";
import { FragmentType, graphql, useFragment } from "../../../libs/gql";
import { FileNodeIcon } from "./FileNodeIcon";

const FileNodeComponent_Fragment = graphql(`
  fragment FileNodeComponent_Fragment on FileNode {
    ...FileNodeIcon_Fragment
    nodeType
    name
    filePath
    offset
    isUpdated
  }
`);

export interface FileNodeComponentProps {
  fragment: FragmentType<typeof FileNodeComponent_Fragment>;
  currentDirectory?: string;
  updateOpenFilePath: (filePath: string) => void;
}

export const FileNodeComponent = (
  props: FileNodeComponentProps
): JSX.Element => {
  const fragment = useFragment(FileNodeComponent_Fragment, props.fragment);
  const offset = fragment.offset ? fragment.offset : 0;
  const isCurrentDirectory =
    fragment.filePath &&
    props.currentDirectory &&
    fragment.filePath === props.currentDirectory;

  const onClick = () => {
    fragment.filePath && props.updateOpenFilePath(fragment.filePath);
  };

  return (
    <div
      css={css`
        display: flex;
        gap: 4px;
        background-color: ${fragment.isUpdated ? "#748d2e" : "#252526"};
        color: white;
        padding-top: 3px;
        padding-bottom: 3px;
        padding-right: 8px;
        padding-left: ${8 * offset + 8}px;
      `}
      onClick={fragment.nodeType === "FILE" ? onClick : undefined}
    >
      <FileNodeIcon fragment={fragment} />
      <div
        css={css`
          background-color: ${isCurrentDirectory && "#007acc"};
          width: fit-content;
          white-space: nowrap;
          font-size: 13px;
          font-family: Menlo, Monaco, Consolas, "Andale Mono", "Ubuntu Mono",
            "Courier New", monospace;
        `}
      >
        {fragment.name}
        {isCurrentDirectory && (
          <span
            css={css`
              color: #a1a1a1;
            `}
          >
            {" //current directory"}
          </span>
        )}
      </div>
    </div>
  );
};
