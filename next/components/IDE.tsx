import { css } from "@emotion/react";
import { IDEEditorTab } from "./IDEEditorTab";
import { IDESideBar } from "./IDESideBar";

export const IDE = (): JSX.Element => {
  return (
    <div
      css={css`
        display: flex;
      `}
    >
      <IDESideBar />
      <IDEEditorTab />
    </div>
  );
};
