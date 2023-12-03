import { FragmentType, graphql, useFragment } from "@/libs/gql";
import styles from "./GqlOpenFilePane.module.css";
import { GqlSourceCodeEditor } from "./editor/GqlSourceCodeEditor";
import { GqlFileNameTabBar } from "./tab/GqlFileNameTabBar";

const fragmentDefinition = graphql(`
  fragment GqlOpenFilePane on OpenFile {
    ...GqlFileNameTabBar
    ...GqlSourceCodeEditor
  }
`);

interface Props {
  fragment: FragmentType<typeof fragmentDefinition>;
  defaultFocusColumn?: string;
}

export function GqlOpenFilePane(props: Props) {
  const fragment = useFragment(fragmentDefinition, props.fragment);

  return (
    <div className={`${styles.component}`}>
      <GqlFileNameTabBar fragment={fragment} />
      <GqlSourceCodeEditor
        fragment={fragment}
        defaultFocusColumn={props.defaultFocusColumn}
      />
    </div>
  );
}
