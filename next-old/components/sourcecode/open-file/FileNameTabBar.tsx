import { css } from "@emotion/react";
import { FileNameTab } from "./FileNameTab";

import { FragmentType, graphql, useFragment } from "../../../libs/gql";

const FileNameTabBar_Fragment = graphql(`
  fragment FileNameTabBar_Fragment on OpenFile {
    ...FileNameTab_Fragment
  }
`);

export interface FileNameTabBarProps {
  fragment: FragmentType<typeof FileNameTabBar_Fragment>;
}

export const FileNameTabBar = (props: FileNameTabBarProps): JSX.Element => {
  const fragment = useFragment(FileNameTabBar_Fragment, props.fragment);
  return (
    <div
      css={css`
        background-color: #222121;
      `}
    >
      <FileNameTab fragment={fragment} />
    </div>
  );
};
