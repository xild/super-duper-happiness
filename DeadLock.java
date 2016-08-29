/**
 * @author Luis Vieira
 */
public class DeadLock {
	public static Object lockedObject1 = new Object();
	public static Object lockedObject2 = new Object();

	public static void main(String args[]) {

		ThreadDeadLock1 T1 = new ThreadDeadLock1();
		ThreadDeadLock2 T2 = new ThreadDeadLock2();
		T2.start();
		T1.start();
	}

	private static class ThreadDeadLock1 extends Thread {
		public void run() {
			synchronized (lockedObject1) {
				String threadName = ThreadDeadLock1.class.getSimpleName();
				System.out.println(threadName + " locking lockedObject 1.");
				try {
					Thread.sleep(10);
				} catch (InterruptedException e) {
					//silence catch. Dont do that in house.
				}
				System.out.println(threadName+" Waiting for lockedObject 1.");
				synchronized (lockedObject2) {
					System.out.println("Now "+threadName +" lock two objects.");
				}
			}
		}
	}

	private static class ThreadDeadLock2 extends Thread {
		public void run() {
			synchronized (lockedObject2) {
				String threadName = ThreadDeadLock2.class.getSimpleName();
				System.out.println(threadName + " Locking lockedObject 2.");
				try {
					Thread.sleep(10);
				} catch (InterruptedException e) {
					// HE-HE
				}
				System.out.println(threadName+" Waiting for lockedObject 1.");
				synchronized (lockedObject1) {
					System.out.println("Now "+threadName +" lock two objects.");
				}
			}
		}
	}
}
