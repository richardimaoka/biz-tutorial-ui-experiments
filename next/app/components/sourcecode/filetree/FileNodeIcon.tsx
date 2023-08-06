import { DirectoryIcon } from "@/app/components/icons/DirectoryIcon";
import { FileIcon } from "@/app/components/icons/FileIcon";

import { FragmentType, graphql, useFragment } from "@/libs/gql";

const fragmentDefinition = graphql(`
  fragment FileNodeIcon_Fragment on FileNode {
    nodeType
  }
`);

export interface FileNodeIconProps {
  fragment: FragmentType<typeof fragmentDefinition>;
}

export const FileNodeIcon = (props: FileNodeIconProps): JSX.Element => {
  const fragment = useFragment(fragmentDefinition, props.fragment);

  if (!fragment.nodeType) {
    return <></>;
  }

  switch (fragment.nodeType) {
    case "FILE":
      return <FileIcon />;
    case "DIRECTORY":
      return <DirectoryIcon />;
    default:
      const _exhaustiveCheck: never = fragment.nodeType;
      return _exhaustiveCheck;
  }
};
