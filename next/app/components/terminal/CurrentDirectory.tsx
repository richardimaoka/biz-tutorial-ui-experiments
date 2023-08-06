import { source_code_pro } from "../fonts/fonts";
import styles from "./style.module.css";
import { DirectoryIcon } from "../icons/DirectoryIcon";

import { FragmentType, graphql, useFragment } from "@/libs/gql";

const fragmentDefinition = graphql(`
  fragment TerminalCurrentDirectory_Fragment on Terminal {
    currentDirectory
  }
`);

interface CurrentDirectoryProps {
  fragment: FragmentType<typeof fragmentDefinition>;
}

export const CurrentDirectory = (props: CurrentDirectoryProps) => {
  const fragment = useFragment(fragmentDefinition, props.fragment);

  return (
    <div className={styles.currentDir}>
      {fragment.currentDirectory && <DirectoryIcon />}
      <span className={source_code_pro.className}>
        {fragment.currentDirectory ? fragment.currentDirectory : "Terminal"}
      </span>
    </div>
  );
};
