import { FileTreeHeader } from "./FileTreeHeader";
import { FileTreeViewer, File } from "./FileTreeViewer";

interface FileTreePaneProps {
  sourceCodeHeight: number;
  files: File[];
}

export const FileTreePane = ({
  files,
  sourceCodeHeight,
}: FileTreePaneProps): JSX.Element => {
  return (
    <div>
      <FileTreeHeader />
      <FileTreeViewer files={files} sourceCodeHeight={sourceCodeHeight} />
    </div>
  );
};
