import { FragmentType, graphql, useFragment } from "@/libs/gql";
import styles from "./style.module.css";

const fragmentDefinition = graphql(`
  fragment FileNameTab_Fragment on OpenFile {
    fileName
  }
`);

export interface FileNameTabProps {
  fragment: FragmentType<typeof fragmentDefinition>;
}

export const FileNameTab = (props: FileNameTabProps): JSX.Element => {
  const fragment = useFragment(fragmentDefinition, props.fragment);
  return <div className={styles.tab}>{fragment.fileName}</div>;
};
