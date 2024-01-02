import { AnglesLeftIcon } from "@/app/components/icons/AnglesLeftIcon";
import { AnglesRightIcon } from "@/app/components/icons/AnglesRightIcon";
import styles from "./GqlFileTreeHeader.module.css";

import { FragmentType, graphql, useFragment } from "@/libs/gql";
import { ProjectDir } from "./ProjectDir";

const fragmentDefinition = graphql(`
  fragment GqlFileTreeHeader on SourceCode {
    projectDir
  }
`);

interface FileTreeHeaderProps {
  fragment: FragmentType<typeof fragmentDefinition>;
  isFolded: boolean;
  onButtonClick: () => void;
}

export const GqlFileTreeHeader = (props: FileTreeHeaderProps): JSX.Element => {
  const fragment = useFragment(fragmentDefinition, props.fragment);

  return (
    <div className={styles.component}>
      <ProjectDir projectDir={fragment.projectDir} />
      <button className={styles.button} onClick={props.onButtonClick}>
        <AnglesLeftIcon />
      </button>
    </div>
  );
};
