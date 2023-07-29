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
  step: string;
}

export const SourceCodeColumn = (props: SourceCodeColumnProps): JSX.Element => {
  const fragment = useFragment(fragmentDefinition, props.fragment);

  if (!fragment.sourceCode) {
    return <></>;
  }

  return <SourceCodeViewer fragment={fragment.sourceCode} step={props.step} />;
};
