import { FragmentType, graphql, useFragment } from "../../../libs/gql";
import { DirectoryIcon } from "./DirectoryIcon";
import { FileIcon } from "./FileIcon";

const FileNodeIcon_Fragment = graphql(`
  fragment FileNodeIcon_Fragment on FileNode {
    nodeType
  }
`);

export interface FileNodeIconProps {
  fragment: FragmentType<typeof FileNodeIcon_Fragment>;
}

export const FileNodeIcon = (props: FileNodeIconProps): JSX.Element => {
  const fragment = useFragment(FileNodeIcon_Fragment, props.fragment);
  if (fragment.nodeType) {
    switch (fragment.nodeType) {
      case "FILE":
        return <FileIcon />;
      case "DIRECTORY":
        return <DirectoryIcon />;
    }
  } else {
    return <></>;
  }

  // switch (fragment.nodeType) {
  //   case default:

  // }
};
