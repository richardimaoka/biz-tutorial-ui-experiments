import { nonNullArray } from "@/libs/nonNullArray";
import { GqlFileNodeComponent } from "../filenode/GqlFileNodeComponent";
import styles from "./style.module.css";

import { FragmentType, graphql, useFragment } from "@/libs/gql";

const fragmentDefinition = graphql(`
  fragment GqlFileTreeComponent on SourceCode2 {
    fileTree {
      filePath
      ...GqlFileNodeComponent
    }
  }
`);

export interface FileTreeComponentProps {
  fragment: FragmentType<typeof fragmentDefinition>;
  isFolded: boolean;
  step: string;
}

export const GqlFileTreeComponent = (
  props: FileTreeComponentProps
): JSX.Element => {
  const fragment = useFragment(fragmentDefinition, props.fragment);

  if (!fragment.fileTree) {
    return <div className={styles.component} />;
  }

  if (props.isFolded) {
    return <div className={styles.component} />;
  }

  const files = nonNullArray(fragment.fileTree);

  return (
    <div className={styles.component}>
      {files.map((file, index) => (
        <GqlFileNodeComponent
          key={file.filePath ? file.filePath : index}
          fragment={file}
          step={props.step}
        />
      ))}
    </div>
  );
};
