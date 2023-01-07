import { FileContentViewer } from "./FileContentViewer";
import { FileNameTabBar } from "./FileNameTabBar";

interface FileContentPaneProps {
  fileContent: string;
  sourceCodeHeight: number;
  prismLanguage: string;
}

export const FileContentPane = ({
  sourceCodeHeight,
  fileContent,
  prismLanguage,
}: FileContentPaneProps): JSX.Element => {
  return (
    <div>
      <FileNameTabBar filename="package.json" />
      <FileContentViewer
        fileContent={fileContent}
        prismLanguage={prismLanguage}
        sourceCodeHeight={sourceCodeHeight}
      />
    </div>
  );
};
