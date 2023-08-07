import { FragmentType, graphql, useFragment } from "@/libs/gql";

import styles from "./style.module.css";
import { FileTreePane } from "./filetree/FileTreePane";
import { FileContentPane } from "./openfile/FileContentPane";

const fragmentDefinition = graphql(`
  fragment SourceCodeColumn_Fragment on SourceCodeColumn {
    sourceCode {
      ...FileTreePane_Fragment
      openFile(filePath: $openFilePath) {
        ...FileContentPane_Fragment
      }
    }
  }
`);

interface SourceCodeColumnProps {
  fragment: FragmentType<typeof fragmentDefinition>;
  step: string;
}

export const SourceCodeColumn = (props: SourceCodeColumnProps) => {
  const fragment = useFragment(fragmentDefinition, props.fragment);

  return (
    <div className={styles.sourcecode}>
      {fragment.sourceCode && (
        <FileTreePane step={props.step} fragment={fragment.sourceCode} />
      )}
      {fragment.sourceCode?.openFile && (
        <FileContentPane fragment={fragment.sourceCode.openFile} />
      )}
    </div>
  );
};
