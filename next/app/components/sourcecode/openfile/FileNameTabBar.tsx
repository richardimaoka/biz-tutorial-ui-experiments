import { FileNameTab } from "./FileNameTab";

import { FragmentType, graphql, useFragment } from "@/libs/gql";
import styles from "./style.module.css";

const FileNameTabBar_Fragment = graphql(`
  fragment FileNameTabBar_Fragment on OpenFile {
    ...FileNameTab_Fragment
  }
`);

export interface FileNameTabBarProps {
  fragment: FragmentType<typeof FileNameTabBar_Fragment>;
}

export const FileNameTabBar = (props: FileNameTabBarProps): JSX.Element => {
  const fragment = useFragment(FileNameTabBar_Fragment, props.fragment);

  return (
    <div className={styles.tabs}>
      <FileNameTab fragment={fragment} />
    </div>
  );
};
