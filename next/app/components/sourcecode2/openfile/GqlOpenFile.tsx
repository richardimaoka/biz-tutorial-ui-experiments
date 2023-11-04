import styles from "./style.module.css";

import { FragmentType, graphql, useFragment } from "@/libs/gql";
import { GqlFileNameTabBar } from "./filetree/header/tab/GqlFileNameTabBar";
import { GqlSourceCodeEditor } from "./filecontent/editor/GqlSourceCodeEditor";

const fragmentDefinition = graphql(`
  fragment GqlFileContentPane on OpenFile {
    ...GqlFileNameTabBar
    ...GqlSourceCodeEditor
  }
`);

interface Props {
  fragment: FragmentType<typeof fragmentDefinition>;
}

export function GqlFileContentPane(props: Props) {
  const fragment = useFragment(fragmentDefinition, props.fragment);

  return (
    <div className={`${styles.component}`}>
      <GqlFileNameTabBar fragment={fragment} />
      <GqlSourceCodeEditor fragment={fragment} />
    </div>
  );
}
