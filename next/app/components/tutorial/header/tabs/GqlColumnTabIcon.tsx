import { SourceCodeIcon } from "@/app/components/icons/SourceCodeIcon";
import { TerminalIcon } from "../../../icons/TerminalIcon";

import { FragmentType, graphql, useFragment } from "@/libs/gql";
import { ChromeIcon } from "@/app/components/icons/ChromeIcon";

const fragmentDefinition = graphql(`
  fragment GqlColumnTabIcon on ColumnWrapper {
    column {
      __typename
    }
  }
`);

interface Props {
  fragment: FragmentType<typeof fragmentDefinition>;
}

export function GqlColumnTabIcon(props: Props): JSX.Element {
  const fragment = useFragment(fragmentDefinition, props.fragment);

  switch (fragment.column.__typename) {
    case "TerminalColumn2":
      return <TerminalIcon />;
    case "SourceCodeColumn2":
      return <SourceCodeIcon />;
    case "BrowserColumn2":
      return <ChromeIcon />;
  }
}
