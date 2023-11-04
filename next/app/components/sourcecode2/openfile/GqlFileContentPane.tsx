import { source_code_pro } from "@/app/components/fonts/fonts";
import styles from "./style.module.css";

import { FragmentType, graphql, useFragment } from "@/libs/gql";
import { GqlFileNameTabBar } from "./tab/GqlFileNameTabBar";
import { GqlSourceCodeEditor } from "./editor/GqlSourceCodeEditor";

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
