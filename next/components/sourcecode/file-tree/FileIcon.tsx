import Image from "next/image";
import FileImg from "../../../public/images/file.png";

export const FileIcon = (): JSX.Element => {
  return <Image src={FileImg} height="16" width="16" alt="file icon" />;
};
