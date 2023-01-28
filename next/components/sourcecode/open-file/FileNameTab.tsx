import { css } from "@emotion/react";
import { FragmentType, graphql, useFragment } from "../../../libs/gql";

const FileNameTab_Fragment = graphql(`
  fragment FileNameTab_Fragment on OpenFile {
    fileName
  }
`);

export interface FileNameTabProps {
  fragment: FragmentType<typeof FileNameTab_Fragment>;
}

export const FileNameTab = (props: FileNameTabProps): JSX.Element => {
  const fragment = useFragment(FileNameTab_Fragment, props.fragment);
  return (
    <div
      css={css`
        width: fit-content;
        font-size: 13px;
        padding: 4px 8px;
        background-color: #232a36;
        color: white;
      `}
    >
      {fragment.fileName}
    </div>
  );
};
