import { FileNodeIcon } from "./FileNodeIcon";
import Link from "next/link";
import styles from "./style.module.css";

import { FragmentType, graphql, useFragment } from "@/libs/gql";

const fragmentDefinition = graphql(`
  fragment FileNodeComponent_Fragment on FileNode {
    ...FileNodeIcon_Fragment
    nodeType
    name
    filePath
    offset
    isUpdated
  }
`);

export interface FileNodeComponentProps {
  fragment: FragmentType<typeof fragmentDefinition>;
  currentDirectory?: string;
  step: string;
}

export const FileNodeComponent = (
  props: FileNodeComponentProps
): JSX.Element => {
  const fragment = useFragment(fragmentDefinition, props.fragment);

  const isCurrentDirectory =
    fragment.filePath &&
    props.currentDirectory &&
    fragment.filePath === props.currentDirectory;

  const filenodeStyle = fragment.isUpdated
    ? `${styles.filenode} ${styles.updated}`
    : styles.filenode;

  const offset = fragment.offset ? fragment.offset : 0;

  const Component = () => (
    <div className={filenodeStyle}>
      <div style={{ marginLeft: `${offset * 16}px` }}>
        <FileNodeIcon fragment={fragment} />
      </div>
      <div className={styles.nodename}>
        {fragment.name}
        {isCurrentDirectory && (
          <span className={styles.currentdirectory}>
            {" //current directory"}
          </span>
        )}
      </div>
    </div>
  );

  return fragment.nodeType === "FILE" && fragment.filePath ? (
    <Link
      style={{
        textDecoration: "none",
      }}
      href={`/?step=${props.step}&openFilePath=${encodeURIComponent(
        fragment.filePath
      )}`}
    >
      <Component />
    </Link>
  ) : (
    <Component />
  );
};
