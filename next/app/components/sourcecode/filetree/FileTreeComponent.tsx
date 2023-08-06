import { nonNullArray } from "@/libs/nonNullArray";
import { FileNodeComponent } from "./FileNodeComponent";
import styles from "./style.module.css";

import { FragmentType, graphql, useFragment } from "@/libs/gql";

const fragmentDefinition = graphql(`
  fragment FileTreeComponent_Fragment on SourceCode {
    fileTree {
      filePath
      ...FileNodeComponent_Fragment
    }
  }
`);

export interface FileTreeComponentProps {
  fragment: FragmentType<typeof fragmentDefinition>;
  isFolded: boolean;
  step: string;
}

export const FileTreeComponent = (
  props: FileTreeComponentProps
): JSX.Element => {
  const fragment = useFragment(fragmentDefinition, props.fragment);

  if (!fragment.fileTree) {
    return <div className={styles.tree} />;
  }

  if (props.isFolded) {
    return <div className={styles.tree} />;
  }

  const files = nonNullArray(fragment.fileTree);

  return (
    <div className={styles.tree}>
      {files.map((file, index) => (
        <FileNodeComponent
          key={file.filePath ? file.filePath : index}
          fragment={file}
          step={props.step}
        />
      ))}
    </div>
  );
};
