import { FileContentViewer } from "./FileContentViewer";
import { FileNameTabBar } from "./FileNameTabBar";
import { FragmentType, graphql, useFragment } from "../../../libs/gql";

const FileContentPane_Fragment = graphql(`
  fragment FileContentPane_Fragment on OpenFile {
    ...FileNameTabBar_Fragment
    ...FileContentViewer_Fragment
  }
`);

export interface FileContentPaneProps {
  fragment: FragmentType<typeof FileContentPane_Fragment>;
}

export const FileContentPane = (props: FileContentPaneProps): JSX.Element => {
  const fragment = useFragment(FileContentPane_Fragment, props.fragment);
  return (
    <div>
      <FileNameTabBar fragment={fragment} />
      <FileContentViewer fragment={fragment} sourceCodeHeight={400} />
    </div>
  );
};
