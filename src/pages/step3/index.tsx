import NextStep from "../components/NextStep";
import { PrintRules } from "../components/PrintRules";

export default function Step3() {
  return (
    <>
      <main className="flex min-h-screen flex-col items-center bg-gradient-to-b from-[#2e026d] to-[#15162c] text-white">
        <div className="mt-10">
          <h1 className="text-4xl">Round 3</h1>
          <div className="mt-10 text-xl">
            <PrintRules label="Deal 10 Cards" />
            <PrintRules label="Two Sequences" />
          </div>
          <div className="mt-10">
            <NextStep label="Step 4" link="/step4" classes="" />
          </div>
        </div>
      </main>
    </>
  );
}
