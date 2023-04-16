import { css } from "@emotion/react";
import { useState } from "react";
import { FragmentType, graphql, useFragment } from "../../../libs/gql";
import { FileTreeComponent } from "./FileTreeComponent";
import { FileTreeHeader } from "./FileTreeHeader";

const FileTreePane_Fragment = graphql(`
  fragment FileTreePane_Fragment on SourceCode {
    ...FileTreeComponent_Fragment
  }
`);

export interface FileTreePaneProps {
  fragment: FragmentType<typeof FileTreePane_Fragment>;
  sourceCodeHeight: number;
  currentDirectory?: string;
  updateOpenFilePath: (filePath: string) => void;
}

export const FileTreePane = (props: FileTreePaneProps): JSX.Element => {
  const fragment = useFragment(FileTreePane_Fragment, props.fragment);
  const [isFolded, setIsFolded] = useState(false);

  return (
    <div>
      <FileTreeHeader
        isFolded={isFolded}
        onButtonClick={() => {
          setIsFolded(!isFolded);
        }}
      />
      <FileTreeComponent
        fragment={fragment}
        sourceCodeHeight={props.sourceCodeHeight}
        isFolded={isFolded}
        currentDirectory={props.currentDirectory}
        updateOpenFilePath={props.updateOpenFilePath}
      />
    </div>
  );
};
