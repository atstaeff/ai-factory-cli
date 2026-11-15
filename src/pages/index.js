import React from 'react';
import Layout from '@theme/Layout';
import Link from '@docusaurus/Link';
import styles from './index.module.css';

const features = [
  {label:'Process-first',title:'Methods stay yours',body:'Keep BMAD, OpenSpec or team-specific methods. Factory is designed to coordinate existing capabilities.'},
  {label:'Brownfield-friendly',title:'Start where you are',body:'Inspect existing project conventions before proposing changes. Read-only scans are available today.'},
  {label:'Enterprise-aware',title:'Approved tools only',body:'The roadmap prioritizes internal CLIs, auditable operations and explicit approval gates.'},
];

export default function Home() {
  return <Layout title="AI Factory" description="Process-driven AI work orchestration">
    <main>
      <header className={styles.hero}>
        <div className="container">
          <div className={styles.eyebrow}>CLI-FIRST · PROCESS-AS-CODE · PROVIDER-NEUTRAL</div>
          <h1>Connect the work.<br/><span className={styles.accent}>Keep the process.</span></h1>
          <p className={styles.lead}>A lightweight, extensible foundation for coordinating humans, AI agents, tools and quality gates without replacing your engineering methods.</p>
          <div className={styles.actions}>
            <Link className="button button--primary button--lg" to="/docs/getting-started">Get started</Link>
            <Link className="button button--secondary button--lg" to="/docs/concepts/architecture">Explore the architecture</Link>
          </div>
          <div className={styles.terminal}>
            <div className={styles.terminalHeader}><span className={styles.dot}/><span className={styles.dot}/><span className={styles.dot}/> factory · terminal</div>
            <pre>{'$ factory scan .\nProject: ./my-project\nPROCESS.md      found    .factory/PROCESS.md\nBMAD            found    _bmad\nOpenSpec        missing  openspec\n\n$ factory processes .\nfeature-development  .factory/PROCESS.md'}</pre>
          </div>
        </div>
      </header>
      <section className={styles.section}>
        <div className="container">
          <h2>Use what is already approved.</h2>
          <p className={styles.sectionLead}>One conceptual model, multiple existing systems. No mandated coding agent, issue tracker or methodology.</p>
          <div className={styles.cards}>
            {features.map(f=><article key={f.label} className={styles.card}><small>{f.label}</small><h3>{f.title}</h3><p>{f.body}</p></article>)}
          </div>
        </div>
      </section>
      <section className={styles.note}><div className="container"><strong>Early prototype.</strong> Scan and discovery commands exist; scheduling, autonomous work, integrations, a live TUI and model routing are future milestones. <Link to="/docs/roadmap">See the roadmap.</Link></div></section>
    </main>
  </Layout>;
}
