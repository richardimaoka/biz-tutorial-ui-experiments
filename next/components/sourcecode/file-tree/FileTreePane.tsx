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
}

export const FileTreePane = (props: FileTreePaneProps): JSX.Element => {
  const fragment = useFragment(FileTreePane_Fragment, props.fragment);

  return (
    <div>
      <FileTreeHeader />
      <FileTreeComponent
        fragment={fragment}
        sourceCodeHeight={props.sourceCodeHeight}
      />
    </div>
  );
};
