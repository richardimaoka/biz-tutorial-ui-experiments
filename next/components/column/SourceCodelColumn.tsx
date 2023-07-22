import { FragmentType, graphql, useFragment } from "../../libs/gql";
import { SourceCodeViewer } from "../sourcecode/SourceCodeViewer";
import { ColumnContentsPosition } from "./ColumnContentsPosition";

const fragmentDefinition = graphql(`
  fragment SourceCodeColumnFragment on SourceCodeColumn {
    sourceCode {
      ...SourceCodeViewer_Fragment
    }
  }
`);

export interface SourceCodeColumnProps {
  fragment: FragmentType<typeof fragmentDefinition>;
}

export const SourceCodeColumn = (props: SourceCodeColumnProps): JSX.Element => {
  const fragment = useFragment(fragmentDefinition, props.fragment);

  console.log("SourceCodeColumn", fragment);

  if (!fragment.sourceCode) {
    return <></>;
  }

  console.log("SourceCodeColumn non-null", fragment);

  return (
    <ColumnContentsPosition position="TOP">
      <SourceCodeViewer fragment={fragment.sourceCode} step="" />
    </ColumnContentsPosition>
  );
};
