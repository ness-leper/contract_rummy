import NextStep from "../components/NextStep";
import { PrintRules } from "../components/PrintRules";

export default function Step2() {
  return (
    <>
      <main className="flex min-h-screen flex-col items-center bg-gradient-to-b from-[#2e026d] to-[#15162c] text-white">
        <div className="mt-10">
          <h1 className="text-4xl">Round 2</h1>
          <div className="mt-10 text-xl">
            <PrintRules label="Deal 10 Cards" />
            <PrintRules label="One Set and One Sequence" />
          </div>
          <div className="mt-10">
            <NextStep label="Step 3" link="/step3" classes="" />
          </div>
        </div>
      </main>
    </>
  );
}
