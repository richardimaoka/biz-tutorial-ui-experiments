import { AnglesLeftIcon } from "@/app/components/icons/AnglesLeftIcon";
import { AnglesRightIcon } from "@/app/components/icons/AnglesRightIcon";
import styles from "./style.module.css";

import { FragmentType, graphql, useFragment } from "@/libs/gql";
import { ProjectDir } from "./ProjectDir";

const fragmentDefinition = graphql(`
  fragment GqlFileTreeHeader on SourceCode2 {
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

  const headerStyle = props.isFolded
    ? `${styles.header} ${styles.folded}`
    : `${styles.header} ${styles.expanded}`;

  return (
    <div className={headerStyle}>
      {!props.isFolded && <ProjectDir projectDir={fragment.projectDir} />}
      <button onClick={props.onButtonClick}>
        {props.isFolded ? <AnglesRightIcon /> : <AnglesLeftIcon />}
      </button>
    </div>
  );
};
