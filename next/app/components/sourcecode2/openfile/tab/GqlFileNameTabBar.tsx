import styles from "./GqlFileNameTabBar.module.css";

import { FileNameTab } from "./FileNameTab";
import { FragmentType, graphql, useFragment } from "@/libs/gql";

const fragmentDefinition = graphql(`
  fragment GqlFileNameTabBar on OpenFile {
    fileName
  }
`);

interface Props {
  fragment: FragmentType<typeof fragmentDefinition>;
}

export function GqlFileNameTabBar(props: Props) {
  const fragment = useFragment(fragmentDefinition, props.fragment);

  if (!fragment.fileName) {
    return <></>;
  }

  return (
    <div className={styles.component}>
      <FileNameTab fileName={fragment.fileName} />
    </div>
  );
}
