import { FragmentType, graphql, useFragment } from "@/libs/gql";
import styles from "./GqlSourceCodeColumn.module.css";
import { GqlOpenFilePane } from "./openfile/GqlOpenFilePane";
import { GqlFileTreePane } from "./filetree/GqlFileTreePane";

const fragmentDefinition = graphql(`
  fragment GqlSourceCodeColumn on SourceCodeColumn {
    sourceCode {
      ...GqlFileTreePane

      openFile(filePath: $file) {
        ...GqlOpenFilePane
      }
    }
  }
`);

interface Props {
  fragment: FragmentType<typeof fragmentDefinition>;
  defaultFocusColumn?: string;
}

export function GqlSourceCodeColumn(props: Props) {
  const fragment = useFragment(fragmentDefinition, props.fragment);

  return (
    <div className={`${styles.component}`}>
      <GqlFileTreePane step="" fragment={fragment.sourceCode} />
      {/* TODO: display an empty open file pane instead of <></> if there is no open */}
      {fragment.sourceCode.openFile && (
        <GqlOpenFilePane
          fragment={fragment.sourceCode.openFile}
          defaultFocusColumn={props.defaultFocusColumn}
        />
      )}
    </div>
  );
}
