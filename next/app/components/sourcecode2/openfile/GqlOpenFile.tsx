import { FragmentType, graphql, useFragment } from "@/libs/gql";
import styles from "./GqlOpenFile.module.css";
import { GqlSourceCodeEditor } from "./editor/GqlSourceCodeEditor";
import { GqlFileNameTabBar } from "./tab/GqlFileNameTabBar";

const fragmentDefinition = graphql(`
  fragment GqlOpenFile on OpenFile {
    ...GqlFileNameTabBar
    ...GqlSourceCodeEditor
  }
`);

interface Props {
  fragment: FragmentType<typeof fragmentDefinition>;
}

export function GqlOpenFile(props: Props) {
  const fragment = useFragment(fragmentDefinition, props.fragment);

  return (
    <div className={`${styles.component}`}>
      <GqlFileNameTabBar fragment={fragment} />
      <GqlSourceCodeEditor fragment={fragment} />
    </div>
  );
}
