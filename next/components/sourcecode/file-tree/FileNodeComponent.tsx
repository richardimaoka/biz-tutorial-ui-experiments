import { css } from "@emotion/react";
import { FragmentType, graphql, useFragment } from "../../../libs/gql";
import { FileNodeIcon } from "./FileNodeIcon";

const FileNodeComponent_Fragment = graphql(`
  fragment FileNodeComponent_Fragment on FileNode {
    ...FileNodeIcon_Fragment
    name
    filePath
    offset
    isUpdated
  }
`);

export interface FileNodeComponentProps {
  fragment: FragmentType<typeof FileNodeComponent_Fragment>;
}

export const FileNodeComponent = (
  props: FileNodeComponentProps
): JSX.Element => {
  const fragment = useFragment(FileNodeComponent_Fragment, props.fragment);
  const offset = fragment.offset ? fragment.offset : 0;
  const background = fragment.isUpdated ? "#748d2e" : "#252526";

  return (
    <div
      css={css`
        display: flex;
        gap: 4px;
        background-color: ${background};
        color: white;
        padding-top: 3px;
        padding-bottom: 3px;
        padding-right: 8px;
        padding-left: ${8 * offset + 8}px;
      `}
    >
      <FileNodeIcon fragment={fragment} />
      <div
        css={css`
          width: fit-content;
          white-space: nowrap;
          font-size: 13px;
          font-family: Menlo, Monaco, Consolas, "Andale Mono", "Ubuntu Mono",
            "Courier New", monospace;
        `}
      >
        {fragment.name}
      </div>
    </div>
  );
};
