import { FileContentViewer } from "./FileContentViewer";
import { FileNameTabBar } from "./FileNameTabBar";
import { FragmentType, graphql, useFragment } from "@/libs/gql";
import { source_code_pro } from "@/app/components/fonts/fonts";
import styles from "./style.module.css";

const fragmentDefinition = graphql(`
  fragment FileContentPane_Fragment on OpenFile {
    ...FileNameTabBar_Fragment
    ...FileContentViewer_Fragment
  }
`);

export interface FileContentPaneProps {
  fragment: FragmentType<typeof fragmentDefinition>;
}

export const FileContentPane = (props: FileContentPaneProps): JSX.Element => {
  const fragment = useFragment(fragmentDefinition, props.fragment);

  return (
    <div className={`${styles.openfile} ${source_code_pro.className}`}>
      <FileNameTabBar fragment={fragment} />
      <FileContentViewer fragment={fragment} />
    </div>
  );
};
