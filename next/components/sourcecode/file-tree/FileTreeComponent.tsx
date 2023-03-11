import { css } from "@emotion/react";

import { FragmentType, graphql, useFragment } from "../../../libs/gql";
import { FileNodeComponent } from "./FileNodeComponent";

const FileTreeComponent_Fragment = graphql(`
  fragment FileTreeComponent_Fragment on SourceCode {
    fileTree {
      filePath
      ...FileNodeComponent_Fragment
    }
  }
`);

export interface FileTreeComponentProps {
  fragment: FragmentType<typeof FileTreeComponent_Fragment>;
  currentDirectory?: string;
  sourceCodeHeight: number;
  isFolded: boolean;
}

export const FileTreeComponent = (
  props: FileTreeComponentProps
): JSX.Element => {
  const fragment = useFragment(FileTreeComponent_Fragment, props.fragment);
  const maxWidth = 160;

  return props.isFolded ? (
    <div
      css={css`
        height: ${props.sourceCodeHeight}px;
        background-color: #252526;
      `}
    />
  ) : (
    <div
      css={css`
        height: ${props.sourceCodeHeight}px;
        max-width: ${maxWidth}px;
        overflow: scroll;
        ::-webkit-scrollbar {
          width: 8px;
          height: 8px;
          background-color: #252526; /* or add it to the track */
        }
        ::-webkit-scrollbar-thumb {
          background: #37373d;
          border-radius: 8px;
        }
        ::-webkit-scrollbar-corner {
          background-color: #252526;
        }
      `}
    >
      <div
        css={css`
          width: fit-content;
          min-width: 100%;
          min-height: 100%; //expand up to the outer element
          background-color: #252526;
        `}
      >
        {fragment.fileTree?.map(
          (elem) =>
            elem && (
              <FileNodeComponent
                key={elem.filePath ? elem.filePath : ""}
                fragment={elem}
                currentDirectory={props.currentDirectory}
              />
            )
        )}
      </div>
    </div>
  );
};
