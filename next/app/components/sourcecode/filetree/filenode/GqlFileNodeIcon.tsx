import { DirectoryIcon } from "@/app/components/icons/DirectoryIcon";
import { FileIcon } from "@/app/components/icons/FileIcon";

import { FragmentType, graphql, useFragment } from "@/libs/gql";

const fragmentDefinition = graphql(`
  fragment GqlFileNodeIcon on FileNode {
    nodeType
  }
`);

interface Props {
  fragment: FragmentType<typeof fragmentDefinition>;
}

export function GqlFileNodeIcon(props: Props): JSX.Element {
  const fragment = useFragment(fragmentDefinition, props.fragment);

  switch (fragment.nodeType) {
    case "FILE":
      return <FileIcon />;
    case "DIRECTORY":
      return <DirectoryIcon />;
  }
}
